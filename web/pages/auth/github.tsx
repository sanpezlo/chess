import { InferGetServerSidePropsType } from "next";

import { APIError } from "@/types/Error";
import { getAuth, withoutAuthSSR } from "@/utils/auth";
import { parseAPIError } from "@/utils/error";

type Props = {
  error?: APIError;
};

export const getServerSideProps = withoutAuthSSR<Props>(async (ctx) => {
  try {
    const { user, setCookie } = await getAuth("github", ctx);
    ctx.res.setHeader("set-cookie", setCookie);
    ctx.req.session.user = user;
    await ctx.req.session.save();
    return {
      redirect: {
        destination: "/dashboard",
        permanent: false,
      },
      props: {},
    };
  } catch (e) {
    return { props: { error: parseAPIError(e) } };
  }
});

export default function Github({
  error,
}: InferGetServerSidePropsType<typeof getServerSideProps>) {
  return (
    <>
      <div>{error?.error}</div>
      <div>{error?.message}</div>
      <div>{error?.suggested}</div>
    </>
  );
}
