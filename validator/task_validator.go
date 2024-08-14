package validator

import (
	"rest-api/model"

	validation "github.com/go-ozzo/ozzo-validation/v4"
)

type ITaskValidator interface {
	TaskValidate(task model.Task) error
}

type taskValidator struct{}

func NewTaskValidator() ITaskValidator {
	return &taskValidator{}
}

func (tv *taskValidator) TaskValidate(task model.Task) error {
	/* taskをバリデーションする
	 */
	return validation.ValidateStruct(&task,
		validation.Field(
			&task.Title,
			validation.Required.Error("title is required"),            // タイトルが入ってるか？
			validation.RuneLength(1, 10).Error("limited max 10 char"), // 文字数が10文字以内か？
		),
	)
}
