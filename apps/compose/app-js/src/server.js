import {
  ResponseOutparam,
  OutgoingBody,
  OutgoingResponse,
  Fields,
} from "wasi:http/types@0.2.3";

import { add } from "example:domain/adder@0.3.0";
import { chat } from "example:domain/chat@0.3.0";

import { Router, json, error, withParams } from "itty-router";

const router = Router({
  before: [withParams], // upstream middleware
  catch: error, // error handling
  finally: [json], // downstream response formatting
});

router
  .get("/hello/:name", ({ name = "World" }) => {
    // simple hello world endpoint
    return `Hello ${name}!`;
  })
  .get("/add", (request) => {
    let { a, b } = request.query;
    return {
      // send the request to the add function, which is the Golang backend component.
      result: add(Number(a), Number(b)),
    };
  })
  .get("/chat", (request) => {
    let prompt = request.query["prompt"];
    let chatRequest = {
      model: "gpt-3.5-turbo",
      messages: [
        {
          role: "user",
          content: prompt,
        },
      ],
    };

    // send the request to the chat function, which is the Golang backend component.
    let result = chat(chatRequest);
    console.log("result", result);
    return {
      result: result,
    };
  });

/**
 * This export represents the `wasi:http/incoming-handler` interface,
 * which describes implementing a HTTP handler in WebAssembly using WASI types.
 *
 * This translates the WASI incoming request into a format that the `itty-router` can understand,
 * then calls the router to handle the request, and finally translates the response back into
 * WASI types.
 */
export const incomingHandler = {
  async handle(incomingRequest, responseOutparam) {
    let requestLike = {
      method: incomingRequest.method().tag.toUpperCase(),
      url: `http://${incomingRequest.authority()}${incomingRequest.pathWithQuery()}`,
      headers: [],
    };

    for (const [key, value] of incomingRequest.headers().entries()) {
      requestLike.headers.push([key, new TextDecoder().decode(value)]);
    }

    let res = await router.fetch(requestLike);

    // Set the status code for the response
    let outgoingResponse = new OutgoingResponse(new Fields());
    outgoingResponse.setStatusCode(res.status);
    // Set the headers for the response
    for (const [key, value] of Object.entries(res.headers)) {
      outgoingResponse
        .headers()
        .set(key, new Uint8Array(new TextEncoder().encode(value)));
    }
    // Finish the response body
    let outgoingBody = outgoingResponse.body();
    {
      // Create a stream for the response body
      let outputStream = outgoingBody.write();

      ensureIterable(res);
      for await (const chunk of res.body) {
        outputStream.blockingWriteAndFlush(chunk);
      }
      outputStream[Symbol.dispose]();
    }

    OutgoingBody.finish(outgoingBody, undefined);
    // Set the created response to an "OK" Result<T> value
    ResponseOutparam.set(outgoingResponse, {
      tag: "ok",
      val: outgoingResponse,
    });
  },
};

function ensureIterable(res) {
  if (!res.body[Symbol.asyncIterator]) {
    res.body[Symbol.asyncIterator] = () => {
      const reader = res.body.getReader();
      return {
        next: () => reader.read(),
      };
    };
  }
}

// handle(incomingRequest, responseOutparam) {
//   // Start building an outgoing response
//   const outgoingResponse = new OutgoingResponse(new Fields());

//   // Access the outgoing response body
//   let outgoingBody = outgoingResponse.body();
//   {
//     // Create a stream for the response body
//     let outputStream = outgoingBody.write();
//     // Write hello world to the response stream
//     let result = add(1 + 2);
//     outputStream.blockingWriteAndFlush(
//       new Uint8Array(
//         new TextEncoder().encode(`Hello from Javascript ${result}!\n`),
//       ),
//     );
//     // @ts-ignore: This is required in order to dispose the stream before we return
//     outputStream[Symbol.dispose]();
//   }

//   // Set the status code for the response
//   outgoingResponse.setStatusCode(200);
//   // Finish the response body
//   OutgoingBody.finish(outgoingBody, undefined);
//   // Set the created response to an "OK" Result<T> value
//   ResponseOutparam.set(outgoingResponse, {
//     tag: "ok",
//     val: outgoingResponse,
//   });
// },
