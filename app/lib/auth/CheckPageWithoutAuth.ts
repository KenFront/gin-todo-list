import { GetServerSideProps } from "next";

import { ROUTE } from "../route";

export const CheckPageWithoutAuth: GetServerSideProps = async (ctx) => {
  const { cookies } = ctx.req;
  if (cookies.auth) {
    return {
      redirect: {
        permanent: false,
        destination: ROUTE.INDEX,
      },
    };
  } else {
    return { props: {} };
  }
};