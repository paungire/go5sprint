package daysteps

import (
	"errors"
	"fmt"
	"strconv"
	"strings"
	"time"

	"github.com/Yandex-Practicum/tracker/internal/personaldata"
	"github.com/Yandex-Practicum/tracker/internal/spentenergy"
)

type DaySteps struct {
	personaldata.Personal
	Steps    int
	Duration time.Duration
}

func (ds *DaySteps) Parse(datastring string) (err error) {
	data := strings.Split(datastring, ",")
	if len(data) != 2 {
		return fmt.Errorf("Некорректные данные: %w", errors.New("Количество параметров не равно двум"))
	}
	steps, err := strconv.Atoi(data[0])
	if err != nil {
		return fmt.Errorf("Некорректные данные: %w", err)
	}
	if steps <= 0 {
		return fmt.Errorf("Некорректные данные: %w", errors.New("Шагов слишком мало"))
	}
	duration, err := time.ParseDuration(data[1])
	if err != nil {
		return fmt.Errorf("Некорректные данные: %w", err)
	}
	if duration <= 0 {
		return fmt.Errorf("Некорректные данные: %w", errors.New("Продолжительность слишком низкая"))
	}
	ds.Steps = steps
	ds.Duration = duration
	return nil
}

func (ds DaySteps) ActionInfo() (string, error) {
	distation := spentenergy.Distance(ds.Steps, ds.Height)
	calories, err := spentenergy.WalkingSpentCalories(ds.Steps, ds.Weight, ds.Height, ds.Duration)
	if err != nil {
		return "", fmt.Errorf("Ошибка вывода: %w", err)
	}
	return fmt.Sprintf(`Количество шагов: %d.
Дистанция составила %.2f км.
Вы сожгли %.2f ккал.`, ds.Steps, distation, calories) + "\n", nil
}
