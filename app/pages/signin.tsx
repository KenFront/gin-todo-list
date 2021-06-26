import { MouseEventHandler, useEffect, useState } from "react";
import {
  Box,
  FormControl,
  FormLabel,
  Input,
  FormErrorMessage,
  Button,
  InputProps,
  InputGroup,
  InputRightElement,
  Stack,
  Spacer,
  Link,
} from "@chakra-ui/react";
import { ViewIcon, ViewOffIcon } from "@chakra-ui/icons";

import { Formik, Form, Field, FormikState } from "formik";

import { CheckPageWithoutAuth } from "@/auth/CheckPageWithoutAuth";
import { signIn } from "@/api/sign";
import { RequestErrorHandler } from "@/lib/request";
import { FullPage } from "@/lib/component/FullPage";
import { Responsive } from "@/lib/component/Responsive";
import { Header } from "@/lib/component/Header";
import { useAppToast } from "@/lib/hook/useAppToast";
import { useAsync } from "@/lib/hook/useAsync";
import { useChangePage } from "@/lib/hook/useChangePage";
import { useRedirectWithAuth } from "@/auth/useRedirectWithAuth";

export const getServerSideProps = CheckPageWithoutAuth;

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

const SignInPage = () => {
  const { changePath } = useChangePage();
  const { toastError } = useAppToast();
  const signInAsync = useAsync(signIn);
  const [showPs, setShowPs] = useState(false);
  const toggleViewPs = () => setShowPs(!showPs);
  const { redirect } = useRedirectWithAuth();

  const registerPath = "/register";
  const toRigisterPage: MouseEventHandler<HTMLAnchorElement> = (e) => {
    e.preventDefault();
    changePath(registerPath);
  };

  useEffect(() => {
    if (signInAsync.status === "success") {
      redirect();
    }
  }, [signInAsync, redirect]);

  useEffect(() => {
    if (signInAsync.status === "error") {
      RequestErrorHandler({
        e: signInAsync.error,
        callback: (str) =>
          toastError({
            title: "Success",
            description: str,
          }),
      });
    }
  }, [signInAsync, toastError]);

  return (
    <FullPage>
      <Header title="Sign in" />
      <Responsive align="center" justify="center">
        <Box w="480px" p={4}>
          <Formik
            initialValues={{ account: "", password: "" }}
            onSubmit={(values) => {
              signInAsync.execute(values);
            }}
          >
            {() => (
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
                      isInvalid={!!form.errors.account && form.touched.account}
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
                        !!form.errors.password && form.touched.password
                      }
                    >
                      <FormLabel htmlFor="password">Password</FormLabel>
                      <InputGroup size="md">
                        <Input
                          {...field}
                          type={showPs ? "text" : "password"}
                          id="password"
                          placeholder="password"
                        />
                        <InputRightElement width="4.5rem">
                          <Button h="1.75rem" size="sm" onClick={toggleViewPs}>
                            {showPs ? <ViewIcon /> : <ViewOffIcon />}
                          </Button>
                        </InputRightElement>
                      </InputGroup>
                      <FormErrorMessage>
                        {form.errors.password}
                      </FormErrorMessage>
                    </FormControl>
                  )}
                </Field>
                <Stack mt={4} direction="row" spacing={4} align="center">
                  <Button
                    colorScheme="teal"
                    isLoading={signInAsync.status === "loading" || signInAsync.status === "success"}
                    type="submit"
                  >
                    Submit
                  </Button>
                  <Spacer />
                  <Link href={registerPath} onClick={toRigisterPage}>
                    Register
                  </Link>
                </Stack>
              </Form>
            )}
          </Formik>
        </Box>
      </Responsive>
    </FullPage>
  );
};

export default SignInPage;
