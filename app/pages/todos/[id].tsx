import { GetServerSideProps } from "next";
import { Heading } from "@chakra-ui/react";

import { CheckPageWithAuth } from "@/auth/CheckPageWithAuth";

import { FullPage } from "@/lib/component/FullPage";
import { Responsive } from "@/lib/component/Responsive";
import { Header } from "@/lib/component/Header";

import { getTodoByIdOnServerSide } from "@/api/todo";
import { GetErrorHandler, UnPromisify } from "@/lib/request";

export const getServerSideProps: GetServerSideProps = async (ctx) => {
  const checkAuth = await CheckPageWithAuth(ctx);
  if (checkAuth.hasOwnProperty("redirect")) {
    return checkAuth;
  }

  const id = ctx.params?.id;

  if (typeof id === "string") {
    try {
      const res = await getTodoByIdOnServerSide({
        id,
        cookie: ctx.req.headers.cookie ?? "",
      });
      return {
        props: {
          res,
        },
      };
    } catch (e) {
      return {
        props: {
          res: e,
        },
      };
    }
  }
  return {
    props: {},
  };
};

const TodoDetailPage = ({
  res,
}: {
  res: UnPromisify<ReturnType<typeof getTodoByIdOnServerSide>>;
}) => {
  const e = GetErrorHandler(res);

  return (
    <FullPage>
      <Header title="Todo Detail" />
      <Responsive align="center" justify="center">
        <Heading size="lg">{e || JSON.stringify(res)}</Heading>
      </Responsive>
    </FullPage>
  );
};

export default TodoDetailPage;
