import { GetServerSideProps } from "next";

export const CheckPageWithAuth: GetServerSideProps = async (ctx) => {
  const { cookies } = ctx.req;
  if (!cookies.auth) {
    return {
      redirect: {
        permanent: false,
        destination: "/signin",
      },
    };
  } else {
    return { props: {} };
  }
};
