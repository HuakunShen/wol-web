import PocketBase from 'pocketbase';
import type { TypedPocketBase } from './pocketbase-types';

export const pb = new PocketBase() as TypedPocketBase;
