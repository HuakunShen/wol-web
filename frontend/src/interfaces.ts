export interface MacInterface {
  id: number;
  user_id: number;
  name: string;
  mac: string;
}

export interface State {
  auth: {
    isAuth: boolean;
    username: string | null;
  };
  computers: Array<MacInterface>;
  errors: Array<string> | null;
  messages: Array<string> | null;
}
