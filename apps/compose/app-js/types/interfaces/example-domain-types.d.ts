/** @module Interface example:domain/types@0.3.0 **/
export interface Message {
  role: string,
  content: string,
}
export interface ChatRequest {
  model: string,
  messages: Array<Message>,
}
export interface Annotation {
  annotationType: string,
  annotationValue: string,
}
export interface Choice {
  index: number,
  message: Message,
  finishReason: string,
}
export interface ChatResponse {
  id: string,
  object: string,
  created: number,
  choices: Array<Choice>,
}
export interface Error {
  code: number,
  message: string,
}
