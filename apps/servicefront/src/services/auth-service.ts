import { AjaxResponse, Post } from "@/lib/axios";

export interface UserResponse {
  id: string;
  email: string;
  first_name: string;
  last_name: string;
  email_verified: boolean;
  role: string;
}

export interface SignUpRequest {
  email: string;
  password: string;
  first_name: string;
  last_name: string;
}

export interface AuthenticationResponse {
  token: string;
  refresh_token: string;
  user: UserResponse;
}

export interface SignInRequest {
  email: string;
  password: string;
}

export class AuthService {
  static async signUp(request: SignUpRequest): AjaxResponse<AuthenticationResponse> {
    return Post<AuthenticationResponse>("auth/signup", request);
  }

  static async signIn(request: SignInRequest): AjaxResponse<AuthenticationResponse> {
    return Post<AuthenticationResponse>("auth/signin", request);
  }
}
