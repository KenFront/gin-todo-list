import {
  FormControl,
  FormLabel,
  Editable,
  EditableInput,
  EditablePreview,
  FormErrorMessage,
  InputProps,
} from "@chakra-ui/react";
import { Field, FormikState, FieldValidator } from "formik";

export const EditableText = ({
  name,
  label,
  placeholder = "",
  validate,
}: {
  name: string;
  label: string;
  placeholder?: string;
  validate: FieldValidator;
}) => {
  return (
    <Field name={name} validate={validate}>
      {({
        field,
        form,
      }: {
        field: Omit<InputProps, "size"> & { size: number };
        form: FormikState<
          {
            [key in string]: string;
          }
        >;
      }) => (
        <FormControl
          mb={4}
          isInvalid={!!form.errors[name] && !!form.touched[name]}
        >
          <FormLabel htmlFor={name}>{label}</FormLabel>
          <Editable defaultValue={`${field.value}`}>
            <EditablePreview />
            <EditableInput {...field} id={name} placeholder={placeholder} />
          </Editable>
          <FormErrorMessage>{form.errors[name]}</FormErrorMessage>
        </FormControl>
      )}
    </Field>
  );
};
