import { GetServerSideProps } from "next";

import { ROUTE } from "../route";

export const CheckPageWithAuth: GetServerSideProps = async (ctx) => {
  const { cookies } = ctx.req;
  if (!cookies.auth) {
    return {
      redirect: {
        permanent: false,
        destination: ROUTE.SIGN_IN,
      },
    };
  } else {
    return { props: {} };
  }
};
