export function getTrackImage(imageUri: string): string {
  return `https://${imageUri.replace("%%", "50x50")}`;
}
