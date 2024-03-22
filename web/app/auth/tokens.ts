import { jwtDecode } from 'jwt-decode';
import { cookies } from 'next/headers';

import { getAccessTokenFromRedisSession, updateRedisSession } from '../lib/redis';

type Tokens = {
  accessToken: string;
  refreshToken: string;
};

type DecodedToken = {
  id: string;
  exp: number;
  iat: number;
};

export async function getNewTokens(refreshToken: string | undefined): Promise<Tokens | null> {
  if (!refreshToken) {
    return null;
  }

  const refreshResponse = await fetch(`${process.env.API_URL}/api/token/refresh`, {
    method: 'POST',
    headers: {
      'Content-Type': 'application/json',
      Cookie: `refresh_token=${refreshToken};` // Send refresh token as a cookie
    },
    credentials: 'include'
  });

  if (refreshResponse.ok) {
    const { accessToken, refreshToken } = await refreshResponse.json();
    return { accessToken, refreshToken };
  }
  return null;
}

export async function getAccessToken(): Promise<string> {
  try {
    const refreshToken = cookies().get('refresh_token')?.value;
    if (!refreshToken) {
      throw new Error('No refresh token found');
    }

    const userId = getUserIdFromRefreshToken(refreshToken as string);
    if (!userId) {
      throw new Error('No user found');
    }

    const accessToken = await getAccessTokenFromRedisSession(userId);

    if (!accessToken) {
      throw new Error('Failed to get new tokens');
    }

    return accessToken;
  } catch (error) {
    console.error('Failed to fetch data:', error);
    throw error;
  }
}

export async function refreshAccessTokenIfNeeded(refreshToken: string): Promise<Tokens> {
  const userId = getUserIdFromRefreshToken(refreshToken);
  if (!userId) throw new Error('No user found');

  const accessToken = await getAccessTokenFromRedisSession(userId);
  if (!accessToken) throw new Error('No access token found');

  const isTokenExpired = checkTokenExpiration(accessToken);
  if (!isTokenExpired) return { accessToken, refreshToken };

  const newTokens = await getNewTokens(refreshToken);
  if (!newTokens) throw new Error('Failed to get new tokens');

  await updateRedisSession(userId, newTokens.accessToken, newTokens.refreshToken);
  return newTokens;
}

export function checkTokenExpiration(accessToken: string): boolean {
  const decoded = jwtDecode<DecodedToken>(accessToken);
  return Date.now() >= decoded.exp * 1000;
}

function getUserIdFromRefreshToken(refreshToken: string): string | null {
  try {
    // Decode the token to access its payload
    const decodedToken = jwtDecode<DecodedToken>(refreshToken);

    const userId = decodedToken.id;

    return userId;
  } catch (error) {
    console.error('Failed to decode refresh token:', error);
    return null;
  }
}
