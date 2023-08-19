package util

import (
	"github.com/bufbuild/protovalidate-go"
	"google.golang.org/genproto/googleapis/rpc/errdetails"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"google.golang.org/protobuf/reflect/protoreflect"
)

func InvalidArgumentError(msg string, violations []*errdetails.BadRequest_FieldViolation) error {
	badRequest := &errdetails.BadRequest{FieldViolations: violations}
	statusInvalid := status.New(codes.InvalidArgument, msg)

	statusDetails, err := statusInvalid.WithDetails(badRequest)
	if err != nil {
		return statusInvalid.Err()
	}

	return statusDetails.Err()
}

func ValidationErrorToFieldViolation(err error) (errs []*errdetails.BadRequest_FieldViolation) {

	if e, ok := err.(*protovalidate.ValidationError); ok {
		for _, ve := range e.Violations {
			errs = append(errs, &errdetails.BadRequest_FieldViolation{
				Field:       ve.FieldPath,
				Description: ve.Message,
			})
		}
	}

	return errs
}

func ValidateRequest(req protoreflect.ProtoMessage) error{
	v, err := protovalidate.New()
	if err != nil {
		return status.Error(codes.Internal, "failed to initialize validator")
	}

	if err = v.Validate(req); err != nil {
		return InvalidArgumentError("validation failed", ValidationErrorToFieldViolation(err))
	}

	return nil
}
