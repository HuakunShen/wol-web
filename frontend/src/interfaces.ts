export interface MacInterface {
  id: Number;
  user_id: Number;
  name: String;
  mac: String;
}
export type Message = {
  id: String;
  message: String;
  variant: String;
};
export interface State {
  auth: {
    isAuth: Boolean;
    username: String | null;
  };
  computers: Array<MacInterface>;
  errors: Array<Message>;
  messages: Array<Message>;
}
