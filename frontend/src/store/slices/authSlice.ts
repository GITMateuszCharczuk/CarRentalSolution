import { createSlice, PayloadAction } from '@reduxjs/toolkit';
import type { UserInfo } from '../../types/api';

interface AuthState {
  user: UserInfo | null;
  token: string | null;
  refresh_token: string | null;
  roles: string[];
  isAuthenticated: boolean;
  isLoading: boolean;
}

interface AuthCredentials {
  user?: UserInfo;
  token: string;
  refresh_token: string;
  roles: string[];
}

// Load initial state from localStorage
const loadState = (): AuthState => {
  try {
    const serializedAuth = localStorage.getItem('auth');
    if (serializedAuth === null) {
      return {
        user: null,
        token: null,
        refresh_token: null,
        roles: [],
        isAuthenticated: false,
        isLoading: false,
      };
    }
    return JSON.parse(serializedAuth);
  } catch (_err) {
    return {
      user: null,
      token: null,
      refresh_token: null,
      roles: [],
      isAuthenticated: false,
      isLoading: false,
    };
  }
};

const initialState: AuthState = loadState();

const authSlice = createSlice({
  name: 'auth',
  initialState,
  reducers: {
    setCredentials: (state, action: PayloadAction<AuthCredentials>) => {
      if (action.payload.user) state.user = action.payload.user;
      state.token = action.payload.token;
      state.refresh_token = action.payload.refresh_token;
      state.roles = action.payload.roles;
      state.isAuthenticated = true;
      state.isLoading = false;
      // Save to localStorage
      localStorage.setItem('auth', JSON.stringify(state));
    },
    setToken: (state, action: PayloadAction<string>) => {
      state.token = action.payload;
      localStorage.setItem('auth', JSON.stringify(state));
    },
    setUser: (state, action: PayloadAction<UserInfo>) => {
      state.user = action.payload;
      localStorage.setItem('auth', JSON.stringify(state));
    },
    setLoading: (state, action: PayloadAction<boolean>) => {
      state.isLoading = action.payload;
    },
    logout: (state) => {
      state.user = null;
      state.token = null;
      state.refresh_token = null;
      state.roles = [];
      state.isAuthenticated = false;
      state.isLoading = false;
      // Clear localStorage
      localStorage.removeItem('auth');
    },
  },
});

// Action creators
export const { setCredentials, setToken, setUser, logout, setLoading } = authSlice.actions;

// Selectors
export const selectCurrentUser = (state: { auth: AuthState }) => state.auth.user;
export const selectIsAuthenticated = (state: { auth: AuthState }) => state.auth.isAuthenticated;
export const selectIsAdmin = (state: { auth: AuthState }) => 
  state.auth.roles.includes('admin');
export const selectAuthToken = (state: { auth: AuthState }) => state.auth.token;
export const selectRefreshToken = (state: { auth: AuthState }) => state.auth.refresh_token;
export const selectIsLoading = (state: { auth: AuthState }) => state.auth.isLoading;
export const selectUserRoles = (state: { auth: AuthState }) => state.auth.roles;

export default authSlice.reducer; 