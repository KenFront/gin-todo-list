import { GetServerSideProps } from "next";

export const CheckPageWithoutAuth: GetServerSideProps = async (ctx) => {
  const { cookies } = ctx.req;
  if (cookies.auth) {
    return {
      redirect: {
        permanent: false,
        destination: "/",
      },
    };
  } else {
    return { props: {} };
  }
};