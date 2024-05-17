export const GC_API = 'https://storage.googleapis.com';
export const BUCKET = process.env.NEXT_PUBLIC_GCP_BUCKET;

export const USER_PROFILE_PIC_URL = `${GC_API}/${BUCKET}/users/profilePics`;
export const USER_COVER_PIC_URL = `${GC_API}/${BUCKET}/users/coverPics`;
export const PROJECT_PIC_URL = `${GC_API}/${BUCKET}/projects`;
export const EVENT_PIC_URL = `${GC_API}/${BUCKET}/events`;
export const POST_PIC_URL = `${GC_API}/${BUCKET}/posts`;
