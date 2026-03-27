package services

import (
	"context"
	"testing"
	"time"

	"github.com/google/uuid"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"gitlab.com/pedrojhrossi/golang-pa/internal/adapters/repository/mocks"
	"gitlab.com/pedrojhrossi/golang-pa/internal/core/domain"
)

func TestAppointmentService_Schedule(t *testing.T) {
	tenantID := uuid.New()
	patientID := uuid.New()

	ctxWithTenant := context.WithValue(context.Background(), "tenant", &domain.Tenant{ID: tenantID})

	now := time.Now().UTC().Truncate(time.Minute)

	tests := []struct {
		name                 string
		startTime            time.Time
		endTime              time.Time
		existingAppointments []*domain.Appointment
		mockRepoErr          error
		wantErr              bool
		expectedErr          error
	}{
		{
			name:                 "Success - No Overlaps",
			startTime:            now.Add(2 * time.Hour),
			endTime:              now.Add(3 * time.Hour),
			existingAppointments: []*domain.Appointment{},
			wantErr:              false,
		},
		{
			name:      "Failure - Conflict Found",
			startTime: now.Add(1 * time.Hour),
			endTime:   now.Add(2 * time.Hour),
			existingAppointments: []*domain.Appointment{
				{
					ID:        uuid.New(),
					StartTime: now.Add(1 * time.Hour),
					EndTime:   now.Add(2 * time.Hour),
				},
			},
			wantErr:     true,
			expectedErr: domain.ErrorOverlap,
		},
		{
			name:        "Failure - Repository Error",
			startTime:   now.Add(1 * time.Hour),
			endTime:     now.Add(2 * time.Hour),
			mockRepoErr: assert.AnError,
			wantErr:     true,
			expectedErr: assert.AnError,
		},
		{
			name:                 "Success - Edge to Edge",
			startTime:            now.Add(1 * time.Hour),
			endTime:              now.Add(2 * time.Hour),
			existingAppointments: []*domain.Appointment{},
			wantErr:              false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			mockRepo := new(mocks.AppointmentRepository)
			service := NewAppointmentService(mockRepo)

			mockRepo.On("FindOverlapping",
				mock.Anything,
				tenantID,
				mock.MatchedBy(func(tm time.Time) bool { return tm.Equal(tt.startTime) }),
				mock.MatchedBy(func(tm time.Time) bool { return tm.Equal(tt.endTime) }),
			).Return(tt.existingAppointments, tt.mockRepoErr)

			if !tt.wantErr {
				mockRepo.On("Create", mock.Anything, mock.MatchedBy(
					func(a *domain.Appointment) bool {
						return a.TenantID == tenantID &&
							a.PatientID == patientID &&
							a.StartTime.Equal(tt.startTime) &&
							a.EndTime.Equal(tt.endTime) &&
							a.ID != uuid.Nil &&
							a.Status == domain.StatusScheduled
					})).Return(nil)
			}

			_, err := service.Schedule(ctxWithTenant, tenantID, patientID, "John Doe", tt.startTime, tt.endTime)

			if tt.wantErr {
				assert.Error(t, err)
				if tt.expectedErr != nil {
					assert.ErrorIs(t, err, tt.expectedErr)
				}
			} else {
				assert.NoError(t, err)
				mockRepo.AssertCalled(t, "Create", mock.Anything, mock.Anything)
			}

			mockRepo.AssertExpectations(t)
		})
	}
}
