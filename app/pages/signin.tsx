import { MouseEventHandler, useEffect } from "react";
import { Box, Button, Stack, Spacer, Link } from "@chakra-ui/react";
import { Formik, Form } from "formik";

import { CheckPageWithoutAuth } from "@/auth/CheckPageWithoutAuth";
import { useRedirectWithAuth } from "@/auth/useRedirectWithAuth";

import { FullPage } from "@/lib/component/FullPage";
import { Responsive } from "@/lib/component/Responsive";
import { Header } from "@/lib/component/Header";
import { TextInput } from "@/lib/component/formik/TextInput";
import { PasswordInput } from "@/lib/component/formik/PasswordInput";

import { useAppToast } from "@/lib/hook/useAppToast";
import { useAsync } from "@/lib/hook/useAsync";
import { useChangePage } from "@/lib/hook/useChangePage";

import { validateAccount } from "@/validator/account";
import { validatePassword } from "@/validator/password";

import { GetErrorHandler } from "@/lib/request";
import { signIn } from "@/api/sign";

import { ROUTE } from "@/route";

export const getServerSideProps = CheckPageWithoutAuth;

const SignInPage = () => {
  const { changePath } = useChangePage();
  const { toastError } = useAppToast();
  const { status, result, error, execute } = useAsync(signIn);
  const { redirect } = useRedirectWithAuth();

  const toRigisterPage: MouseEventHandler<HTMLAnchorElement> = (e) => {
    e.preventDefault();
    changePath({ path: ROUTE.REGISTER });
  };

  useEffect(() => {
    if (status === "success" && result) {
      redirect();
    }
  }, [status, result, redirect]);

  useEffect(() => {
    if (status === "error" && error) {
      const e = GetErrorHandler(error);
      toastError({
        title: "Error",
        description: e,
      });
    }
  }, [status, error, toastError]);

  return (
    <FullPage>
      <Header title="Sign in" />
      <Responsive align="center" justify="center">
        <Box w="480px" p={4}>
          <Formik
            initialValues={{ account: "", password: "" }}
            onSubmit={(values) => {
              execute(values);
            }}
          >
            <Form>
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
              <Stack mt={4} direction="row" spacing={4} align="center">
                <Button
                  colorScheme="teal"
                  isLoading={status === "loading" || status === "success"}
                  type="submit"
                >
                  Submit
                </Button>
                <Spacer />
                <Link href={ROUTE.REGISTER} onClick={toRigisterPage}>
                  Register
                </Link>
              </Stack>
            </Form>
          </Formik>
        </Box>
      </Responsive>
    </FullPage>
  );
};

export default SignInPage;
