import { useEffect, useMemo } from "react";
import { Table, Column } from "react-virtualized";
import { Checkbox, HStack, Button, Stack, Skeleton } from "@chakra-ui/react";
import { format } from "date-fns";
import { EditIcon } from "@chakra-ui/icons";

import { CheckPageWithAuth } from "@/auth/CheckPageWithAuth";
import { SignOutButton } from "@/auth/SignOutButton";
import { FullPage } from "@/lib/component/FullPage";
import { Header } from "@/lib/component/Header";
import { Responsive } from "@/lib/component/Responsive";

import { useAsync } from "@/lib/hook/useAsync";
import { useChangePage } from "@/lib/hook/useChangePage";

import { getTodos } from "@/api/todo";

import { ROUTE } from "@/route";

export const getServerSideProps = CheckPageWithAuth;

const TodolistPage = () => {
  const { execute, result } = useAsync(getTodos);
  const { changePath } = useChangePage();

  const toTodoAddPage = () => {
    changePath(ROUTE.TODO_ADD);
  };

  const colW = useMemo(() => [100, 200, 100, 250, 250], []);
  const colSum = colW.reduce((sum, num) => sum + num);

  useEffect(() => {
    execute();
  }, [execute]);

  return (
    <FullPage>
      <Header title="Todo list" rightArea={<SignOutButton />} />
      <Responsive p={4} justify="center">
        <HStack w={colSum} spacing={4}>
          <Button
            leftIcon={<EditIcon />}
            colorScheme="facebook"
            variant="solid"
            onClick={toTodoAddPage}
          >
            Add
          </Button>
        </HStack>
        {result ? (
          <Table
            width={colSum}
            height={360}
            headerHeight={28}
            rowHeight={60}
            rowCount={result.data.length}
            rowGetter={({ index }) => result.data[index]}
          >
            <Column width={colW[0]} label="Title" dataKey="title" />
            <Column width={colW[1]} label="Description" dataKey="description" />
            <Column
              width={colW[2]}
              label="Status"
              dataKey="status"
              cellRenderer={({ cellData }) => (
                <Checkbox isChecked={cellData === "completed"} isDisabled />
              )}
            />
            <Column
              width={colW[3]}
              label="CreatedAt"
              dataKey="createdAt"
              cellRenderer={({ cellData }) => (
                <span>
                  {format(new Date(cellData), "yyyy-MM-dd HH:mm:ss z")}
                </span>
              )}
            />
            <Column
              width={colW[4]}
              label="UpdatedAt"
              dataKey="updatedAt"
              cellRenderer={({ cellData }) => (
                <span>
                  {format(new Date(cellData), "yyyy-MM-dd HH:mm:ss z")}
                </span>
              )}
            />
          </Table>
        ) : (
          <Stack w={colSum}>
            {new Array(18).fill("").map((_val, i) => (
              <Skeleton key={i} height="20px" />
            ))}
          </Stack>
        )}
      </Responsive>
    </FullPage>
  );
};

export default TodolistPage;
