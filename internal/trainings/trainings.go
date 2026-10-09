package trainings

import (
	"errors"
	"fmt"
	"strconv"
	"strings"
	"time"

	"github.com/Yandex-Practicum/tracker/internal/personaldata"
	"github.com/Yandex-Practicum/tracker/internal/spentenergy"
)

type Training struct {
	personaldata.Personal
	Steps        int
	TrainingType string
	Duration     time.Duration
}

func (t *Training) Parse(datastring string) (err error) {
	data := strings.Split(datastring, ",")
	if len(data) != 3 {
		return fmt.Errorf("Некорректные данные: %w", errors.New("Количество параметров не равно трём"))
	}
	steps, err := strconv.Atoi(data[0])
	if err != nil {
		return fmt.Errorf("Некорректные данные: %w", err)
	}
	if steps <= 0 {
		return fmt.Errorf("Некорректные данные: %w", errors.New("Шагов слишком мало"))
	}

	duration, err := time.ParseDuration(data[2])
	if err != nil {
		return fmt.Errorf("Некорректные данные: %w", err)
	}
	if duration <= 0 {
		return fmt.Errorf("Некорректные данные: %w", errors.New("Продолжительность слишком низкая"))
	}

	t.Steps = steps
	t.TrainingType = data[1]
	t.Duration = duration
	return nil
}

func (t Training) ActionInfo() (string, error) {
	distation := spentenergy.Distance(t.Steps, t.Height)
	meanSpeed := spentenergy.MeanSpeed(t.Steps, t.Height, t.Duration)

	var calories float64
	var err error
	switch t.TrainingType {
	case "Ходьба":
		calories, err = spentenergy.WalkingSpentCalories(t.Steps, t.Weight, t.Height, t.Duration)
	case "Бег":
		calories, err = spentenergy.RunningSpentCalories(t.Steps, t.Weight, t.Height, t.Duration)
	default:
		err = errors.New("Неизвестный тип тренировки")
	}
	if err != nil {
		return "", fmt.Errorf("Ошибка расчета калорий: %w", err)
	}

	return fmt.Sprintf(`Тип тренировки: %s
Длительность: %.2f ч.
Дистанция: %.2f км.
Скорость: %.2f км/ч
Сожгли калорий: %.2f`,
		t.TrainingType, t.Duration.Hours(), distation, meanSpeed, calories) + "\n", nil
}
