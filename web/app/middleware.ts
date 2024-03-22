import { NextRequest, NextResponse } from 'next/server';

import { refreshAccessTokenIfNeeded } from './auth/tokens';

export async function middleware(req: NextRequest) {
  const refreshToken = req.cookies.get('refresh_token');
  if (!refreshToken) return NextResponse.redirect(new URL('/login', req.url));

  const newTokens = await refreshAccessTokenIfNeeded(refreshToken.value);
  if (newTokens) {
    // Update cookies and headers as necessary
    const response = NextResponse.next();
    response.cookies.set('refresh_token', newTokens.refreshToken, {
      httpOnly: true,
      sameSite: 'strict',
      secure: true
    });
    return response;
  }

  return NextResponse.redirect(new URL('/login', req.url));
}
