const defaultConfig: RequestInit = {
  cache: "no-cache",
  credentials: "same-origin",
  headers: {
    "content-type": "application/json",
  },
  redirect: "follow",
  referrer: "no-referrer",
};
Object.freeze(defaultConfig);

export const Request = async ({
  path,
  options,
}: {
  path: string;
  options: RequestInit;
}) => {
  const response = await fetch(path, {
    ...defaultConfig,
    ...options,
  });

  let jsonRes: unknown;

  try {
    jsonRes = await response.json();
  } catch (e) {
    throw {
      error: "Something wrong",
    };
  }

  if (response.ok) {
    return jsonRes;
  } else {
    throw jsonRes;
  }
};

export const RequestErrorHandler = ({
  e,
  callback,
}: {
  // eslint-disable-next-line @typescript-eslint/no-explicit-any
  e: any;
  callback?: (e: string) => void;
}) => {
  if (e && e.error && typeof e.error === "string") {
    console.error(e);
    callback && callback(e.error);
  }
};

export const PostOptions = (body?: unknown): RequestInit => ({
  method: "post",
  body: body ? JSON.stringify(body) : null,
});
