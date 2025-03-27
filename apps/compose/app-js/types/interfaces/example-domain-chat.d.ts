/** @module Interface example:domain/chat@0.3.0 **/
export function chat(request: ChatRequest): ChatResponse;
export type ChatRequest = import('./example-domain-types.js').ChatRequest;
export type ChatResponse = import('./example-domain-types.js').ChatResponse;
export type Error = import('./example-domain-types.js').Error;
