import PocketBase from 'pocketbase';
import type { TypedPocketBase } from './pocketbase-types';
import { dev } from '$app/environment';

export const pb = new PocketBase(dev ? 'http://localhost:8090' : undefined) as TypedPocketBase;
