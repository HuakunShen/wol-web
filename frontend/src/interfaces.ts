export interface MacInterface {
  id: number;
  user_id: number;
  name: string;
  mac: string;
}
export type Message = {
  id: string;
  message: string;
  variant: string;
};
export interface State {
  auth: {
    isAuth: boolean;
    username: string | null;
  };
  computers: Array<MacInterface>;
  errors: Array<Message>;
  messages: Array<Message>;
}
