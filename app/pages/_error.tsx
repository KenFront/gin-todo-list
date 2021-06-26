import { GetServerSideProps } from "next";
import { Heading } from "@chakra-ui/react";

import { FullPage } from "@/lib/component/FullPage";
import { Responsive } from "@/lib/component/Responsive";
import { Header } from "@/lib/component/Header";

export const getServerSideProps: GetServerSideProps = async ({ res }) => {
  return {
    props: {
      statusCode: res.statusCode,
    },
  };
};

const ErrorMsg = (code: number) => {
  switch (true) {
    case code >= 500:
      return "Server is not working now";
    case code === 404:
      return "Page not found";
    case code >= 400:
      return "Something wrong";
  }
};

const ErrorPage = ({ statusCode }: { statusCode: number }) => {
  return (
    <FullPage>
      <Header title="Error" />
      <Responsive align="center" justify="center">
        <Heading size="lg">{ErrorMsg(statusCode)}</Heading>
      </Responsive>
    </FullPage>
  );
};

export default ErrorPage;
