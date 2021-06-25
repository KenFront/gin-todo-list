import { GetServerSideProps } from "next";
import {
  Flex,
  Box,
  FormControl,
  FormLabel,
  Input,
  FormErrorMessage,
  Button,
  InputProps,
  useToast,
} from "@chakra-ui/react";
import { Formik, Form, Field, FormikState } from "Formik";

export const getServerSideProps: GetServerSideProps = async (ctx) => {
  const { req } = ctx;

  const { cookies } = req;

  if (!!cookies.auth) {
    return { props: { status: "authorized" } };
  } else {
    return { props: { status: "unauthorized" } };
  }
};

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

const signOut = async () => {
  const res = await fetch("/api/signout", {
    method: "POST",
  }).then((res) => res.json());
  return res;
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

const SignInForm = () => {
  const toast = useToast();

  return (
    <Flex w="100vw" h="100vh" align="center" justify="center">
      <Box w="480px" p={4}>
        <Formik
          initialValues={{ account: "", password: "" }}
          onSubmit={async (values, actions) => {
            try {
              const res = await signIn(values);
              toast({
                position: "top",
                title: "Success",
                description: res.data,
                status: "success",
                duration: 3000,
                isClosable: true,
                onCloseComplete: () => {
                  window.location.reload();
                },
              });
            } catch (error) {
              console.error(error);
              toast({
                position: "top",
                title: "Error",
                description: error,
                status: "error",
                duration: 3000,
                isClosable: true,
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
    </Flex>
  );
};

const TestingPage = () => {
  const toast = useToast();

  return (
    <Flex w="100vw" h="100vh" align="center" justify="center">
      <Button
        mt={4}
        colorScheme="teal"
        onClick={async () => {
          const res = await signOut();
          toast({
            position: "top",
            title: "Success",
            description: res.data,
            status: "success",
            duration: 3000,
            isClosable: true,
            onCloseComplete: () => {
              window.location.reload();
            },
          });
        }}
      >
        Sign out
      </Button>
    </Flex>
  );
};

const IndexPage = ({ status }: { status: "authorized" | "unauthorized" }) => {
  switch (status) {
    case "authorized":
      return <TestingPage />;
    case "unauthorized":
      return <SignInForm />;
  }
};

export default IndexPage;
