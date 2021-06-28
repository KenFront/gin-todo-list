import {
  FormControl,
  FormLabel,
  Text as ChakraText,
  InputProps,
} from "@chakra-ui/react";
import { Field } from "formik";

export const Text = ({ name, label }: { name: string; label: string }) => {
  return (
    <Field name={name}>
      {({
        field,
      }: {
        field: InputProps;
      }) => (
        <FormControl mb={4}>
          <FormLabel>{label}</FormLabel>
          <ChakraText fontSize="md">{field.value}</ChakraText>
        </FormControl>
      )}
    </Field>
  );
};
