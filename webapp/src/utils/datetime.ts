export const formatDuration = (durationInMs: number): string => {
  const totalSeconds = Math.ceil(durationInMs / 1000);
  const minutes = Math.ceil(totalSeconds / 60);
  const seconds = totalSeconds % 60;

  const minutesStr = String(minutes).padStart(2, '0');
  const secondsStr = String(seconds).padStart(2, '0');

  return `${minutesStr}:${secondsStr}`;
};
