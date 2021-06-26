import { useState } from "react";
import {
  FormControl,
  FormLabel,
  Input,
  FormErrorMessage,
  InputProps,
  InputGroup,
  InputRightElement,
  Button,
} from "@chakra-ui/react";
import { ViewIcon, ViewOffIcon } from "@chakra-ui/icons";
import { Field, FormikState, FieldValidator } from "formik";

export const PasswordInput = <
  T extends {
    [key in string]: string;
  }
>({
  name,
  validate,
  placeholder = "",
}: {
  name: string;
  placeholder?: string;
  validate: FieldValidator;
}) => {
  const [showPs, setShowPs] = useState(false);
  const toggleViewPs = () => setShowPs(!showPs);
  return (
    <Field name={name} validate={validate}>
      {({ field, form }: { field: InputProps; form: FormikState<T> }) => (
        <FormControl isInvalid={!!form.errors[name] && !!form.touched[name]}>
          <FormLabel mt={4} htmlFor={name}>
            Password
          </FormLabel>
          <InputGroup size="md">
            <Input
              {...field}
              type={showPs ? "text" : "password"}
              id={name}
              placeholder={placeholder}
            />
            <InputRightElement width="4.5rem">
              <Button h="1.75rem" size="sm" onClick={toggleViewPs}>
                {showPs ? <ViewIcon /> : <ViewOffIcon />}
              </Button>
            </InputRightElement>
          </InputGroup>
          <FormErrorMessage>{form.errors[name]}</FormErrorMessage>
        </FormControl>
      )}
    </Field>
  );
};
