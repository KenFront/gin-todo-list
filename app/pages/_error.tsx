import { GetServerSideProps } from "next";
import { Heading } from "@chakra-ui/react";

import { FullPage } from "@/lib/component/FullPage";

export const getServerSideProps: GetServerSideProps = async ({ res }) => {
  return {
    props: {
      statusCode: res.statusCode,
    },
  };
};

const ErrorPage = ({ statusCode }: { statusCode: number }) => {
  switch (true) {
    case statusCode >= 500:
      return (
        <FullPage>
          <Heading size="lg">Server is not working now</Heading>
        </FullPage>
      );
    case statusCode === 404:
      return (
        <FullPage>
          <Heading size="lg">Page not found</Heading>
        </FullPage>
      );
    case statusCode >= 400:
      return (
        <FullPage>
          <Heading size="lg">Something wrong</Heading>
        </FullPage>
      );
  }
};

export default ErrorPage;
