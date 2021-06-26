import {
  FormControl,
  FormLabel,
  Input,
  FormErrorMessage,
  InputProps,
} from "@chakra-ui/react";
import { Field, FormikState, FieldValidator } from "formik";

export const TextInput = ({
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
        field: InputProps;
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
          <Input {...field} id={name} placeholder={placeholder} />
          <FormErrorMessage>{form.errors[name]}</FormErrorMessage>
        </FormControl>
      )}
    </Field>
  );
};
