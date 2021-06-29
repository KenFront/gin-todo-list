import {
  FormControl,
  FormLabel,
  Switch as ChakraSwitch,
  InputProps,
} from "@chakra-ui/react";
import { Field } from "formik";

export const Switch = ({ name, label }: { name: string; label: string }) => {
  return (
    <Field name={name}>
      {({
        field,
      }: {
        field: InputProps;
      }) => (
        <FormControl mb={4}>
          <FormLabel htmlFor={name}>{label}</FormLabel>
          <ChakraSwitch isChecked={!!field.value} onChange={field.onChange} id={name} isFocusable />
        </FormControl>
      )}
    </Field>
  );
};
