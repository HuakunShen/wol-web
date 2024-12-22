import { $ } from 'bun';
import fs from 'fs';

// copy build folder to ../server/pb_public
await $`rm -rf ../server/pb_public`;
fs.cpSync('build', '../server/pb_public', { recursive: true });
