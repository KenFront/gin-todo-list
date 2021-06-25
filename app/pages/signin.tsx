import {
  Box,
  FormControl,
  FormLabel,
  Input,
  FormErrorMessage,
  Button,
  InputProps,
} from "@chakra-ui/react";
import { Formik, Form, Field, FormikState } from "Formik";

import { CheckPageWithoutAuth } from "@/lib/auth/CheckPageWithoutAuth";
import { FullPage } from "@/lib/component/FullPage";
import { useAppToast } from "@/lib/hook/useAppToast";

export const getServerSideProps = CheckPageWithoutAuth;

const signIn = async (val: { account: string; password: string }) => {
  try {
    const res = await fetch("/api/signin", {
      body: JSON.stringify(val),
      method: "POST",
    }).then((res) => res.json());
    if (!!res.error) {
      throw res.error;
    }
    return res;
  } catch (error) {
    throw error;
  }
};

function validateSignIn({ name, value }: { name: string; value: string }) {
  switch (true) {
    case !value:
      return `${name} is required`;
    case value.length < 3:
      return `${name} is too short`;
    case value.length > 100:
      return `${name} is too long`;
    default:
      return "";
  }
}

function validateAccount() {
  return (value: string) => validateSignIn({ name: "Acconut", value });
}

function validatePassword() {
  return (value: string) => validateSignIn({ name: "password", value });
}

const IndexPage = () => {
  const { toastSuccess, toastError } = useAppToast();

  return (
    <FullPage>
      <Box w="480px" p={4}>
        <Formik
          initialValues={{ account: "", password: "" }}
          onSubmit={async (values, actions) => {
            try {
              const res = await signIn(values);
              toastSuccess({
                title: "Success",
                description: res.data,
                onCloseComplete: () => {
                  window.location.reload();
                },
              });
            } catch (error) {
              console.error(error);
              toastError({
                title: "Error",
                description: error,
              });
            } finally {
              actions.setSubmitting(false);
            }
          }}
        >
          {(props) => (
            <Form>
              <Field name="account" validate={validateAccount()}>
                {({
                  field,
                  form,
                }: {
                  field: InputProps;
                  form: FormikState<{ account: string }>;
                }) => (
                  <FormControl
                    isInvalid={!!form.errors.account && !!form.touched.account}
                  >
                    <FormLabel htmlFor="account">Account</FormLabel>
                    <Input {...field} id="account" placeholder="account" />
                    <FormErrorMessage>{form.errors.account}</FormErrorMessage>
                  </FormControl>
                )}
              </Field>
              <Field name="password" validate={validatePassword()}>
                {({
                  field,
                  form,
                }: {
                  field: InputProps;
                  form: FormikState<{ password: string }>;
                }) => (
                  <FormControl
                    isInvalid={
                      !!form.errors.password && !!form.touched.password
                    }
                  >
                    <FormLabel htmlFor="password">Password</FormLabel>
                    <Input {...field} id="password" placeholder="password" />
                    <FormErrorMessage>{form.errors.password}</FormErrorMessage>
                  </FormControl>
                )}
              </Field>
              <Button
                mt={4}
                colorScheme="teal"
                isLoading={props.isSubmitting}
                type="submit"
              >
                Submit
              </Button>
            </Form>
          )}
        </Formik>
      </Box>
    </FullPage>
  );
};

export default IndexPage;
