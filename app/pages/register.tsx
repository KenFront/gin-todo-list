import { useEffect } from "react";
import { Box, Button, Stack } from "@chakra-ui/react";
import { Formik, Form } from "formik";

import { FullPage } from "@/lib/component/FullPage";
import { Header } from "@/lib/component/Header";
import { Responsive } from "@/lib/component/Responsive";
import { TextInput } from "@/lib/component/formik/TextInput";
import { PasswordInput } from "@/lib/component/formik/PasswordInput";

import { useAsync } from "@/lib/hook/useAsync";
import { useAppToast } from "@/lib/hook/useAppToast";
import { useChangePage } from "@/lib/hook/useChangePage";

import { validateAccount } from "@/validator/account";
import { validatePassword } from "@/validator/password";
import { validateConfirmPassword } from "@/validator/confirmPassword";
import { validateEmail } from "@/validator/email";
import { isNotEmpty } from "@/validator/isNotEmpty";

import { RequestErrorHandler } from "@/lib/request";
import { register } from "@/api/user";

import { ROUTE } from "@/route";

const IndexPage = () => {
  const registerAsync = useAsync(register);
  const { toastSuccess, toastError } = useAppToast();
  const { changePath } = useChangePage();

  useEffect(() => {
    if (registerAsync.status === "success") {
      toastSuccess({
        title: "Success",
        description: "Register successfully",
        onCloseComplete: () => {
          changePath(ROUTE.SIGN_IN);
        },
      });
      registerAsync.reset();
    }
  }, [registerAsync, changePath, toastSuccess]);
  useEffect(() => {
    if (registerAsync.status === "error") {
      RequestErrorHandler({
        e: registerAsync.error,
        callback: (str) =>
          toastError({
            title: "Error",
            description: str,
          }),
      });
      registerAsync.reset();
    }
  }, [registerAsync, toastError]);

  return (
    <FullPage>
      <Header title="Register" />
      <Responsive align="center" justify="center">
        <Box w="480px" p={4}>
          <Formik
            initialValues={{
              name: "",
              account: "",
              password: "",
              confirmPassword: "",
              email: "",
            }}
            onSubmit={(values) => {
              registerAsync.execute({
                name: values.name,
                account: values.account,
                password: values.password,
                email: values.email,
              });
            }}
          >
            {(props) => (
              <Form>
                <TextInput
                  name="name"
                  label="Name"
                  validate={isNotEmpty("Name")}
                  placeholder="name"
                />
                <TextInput
                  name="account"
                  label="Account"
                  validate={validateAccount}
                  placeholder="account"
                />
                <PasswordInput
                  name="password"
                  label="Password"
                  validate={validatePassword}
                  placeholder="password"
                />
                <PasswordInput
                  name="confirmPassword"
                  label="Confirm Password"
                  validate={validateConfirmPassword(props.values.password)}
                  placeholder="confirmPassword"
                />
                <TextInput
                  name="email"
                  label="Email"
                  validate={validateEmail}
                  placeholder="email"
                />
                <Stack
                  mt={4}
                  direction="row"
                  spacing={4}
                  justify="center"
                  align="center"
                >
                  <Button
                    colorScheme="teal"
                    isLoading={
                      registerAsync.status === "loading" ||
                      registerAsync.status === "success"
                    }
                    type="submit"
                  >
                    Submit
                  </Button>
                </Stack>
              </Form>
            )}
          </Formik>
        </Box>
      </Responsive>
    </FullPage>
  );
};

export default IndexPage;
