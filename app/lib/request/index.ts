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

export type UnPromisify<T> = T extends Promise<infer U> ? U : T;


export const Request = async ({
  path,
  options,
}: {
  path: string;
  options?: RequestInit;
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

// eslint-disable-next-line @typescript-eslint/no-explicit-any
export const GetErrorHandler = (e: any): string => {
  if (e && e.error && typeof e.error === "string") {
    console.error(e);
    return e.error;
  } else {
    return "";
  }
};

export const PostOptions = (body?: unknown): RequestInit => ({
  method: "post",
  body: body ? JSON.stringify(body) : null,
});
