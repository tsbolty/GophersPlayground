import Redis from 'ioredis';

export const redis = new Redis({
  host: process.env.REDIS_HOST || 'localhost',
  port: Number(process.env.REDIS_PORT) || 6379,
  password: process.env.REDIS_PASSWORD || '',
  db: Number(process.env.REDIS_DB) || 0
});

export async function updateRedisSession(userId: string, accessToken: string, refreshToken: string): Promise<void> {
  const sessionKey = `session:${userId}`;
  const sessionData = JSON.stringify({ accessToken, refreshToken });
  await redis.set(sessionKey, sessionData, 'EX', 24 * 60 * 60); // Set with a 24-hour expiration
}

export async function getAccessTokenFromRedisSession(userId: string): Promise<string> {
  const sessionKey = `session:${userId}`;
  const sessionData = await redis.get(sessionKey);
  if (!sessionData) {
    throw new Error('No session found');
  }
  return JSON.parse(sessionData).accessToken;
}
