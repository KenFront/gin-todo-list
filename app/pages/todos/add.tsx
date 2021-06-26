import { useEffect } from "react";
import { Box, Button, Stack } from "@chakra-ui/react";
import { Formik, Form } from "formik";

import { CheckPageWithAuth } from "@/auth/CheckPageWithAuth";
import { SignOutButton } from "@/auth/SignOutButton";
import { FullPage } from "@/lib/component/FullPage";
import { Header } from "@/lib/component/Header";
import { Responsive } from "@/lib/component/Responsive";
import { TextInput } from "@/lib/component/formik/TextInput";

import { useAsync } from "@/lib/hook/useAsync";
import { useAppToast } from "@/lib/hook/useAppToast";

import { RequestErrorHandler } from "@/lib/request";
import { addTodo } from "@/api/todo";

import { isNotEmpty } from "@/validator/isNotEmpty";

export const getServerSideProps = CheckPageWithAuth;

const TodoAddPage = () => {
  const addTodoAsync = useAsync(addTodo);
  const { toastSuccess, toastError } = useAppToast();

  useEffect(() => {
    if (addTodoAsync.status === "success") {
      toastSuccess({
        title: "Success",
        description: "Add todo successfully",
      });
      addTodoAsync.reset();
    }
  }, [addTodoAsync, toastSuccess]);

  useEffect(() => {
    if (addTodoAsync.status === "error") {
      RequestErrorHandler({
        e: addTodoAsync.error,
        callback: (str) =>
          toastError({
            title: "Error",
            description: str,
          }),
      });
      addTodoAsync.reset();
    }
  }, [addTodoAsync, toastError]);

  return (
    <FullPage>
      <Header title="Todo Add" rightArea={<SignOutButton />} />
      <Responsive align="center" justify="center">
        <Box w="480px" p={4}>
          <Formik
            initialValues={{ title: "", description: "" }}
            onSubmit={async(values, action) => {
              await addTodoAsync.execute(values);
              action.resetForm();
            }}
          >
            <Form>
              <TextInput
                name="title"
                label="Title"
                validate={isNotEmpty("Title")}
                placeholder="title"
              />
              <TextInput
                name="description"
                label="Description"
                validate={isNotEmpty("Description")}
                placeholder="description"
              />
              <Stack mt={4} direction="row" spacing={4} align="center">
                <Button
                  colorScheme="teal"
                  isLoading={
                    addTodoAsync.status === "loading" ||
                    addTodoAsync.status === "success"
                  }
                  type="submit"
                >
                  Submit
                </Button>
              </Stack>
            </Form>
          </Formik>
        </Box>
      </Responsive>
    </FullPage>
  );
};

export default TodoAddPage;
