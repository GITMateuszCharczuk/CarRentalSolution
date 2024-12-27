export const formatDateForApi = (date: string | Date): string => {
  const d = new Date(date);
  // Format in RFC3339 format (e.g., "2023-12-26T15:30:00Z" or "2023-12-26T15:30:00+01:00")
  return d.toISOString();
};

// Convert API datetime string to local datetime-local input format
export const formatDateForInput = (dateStr: string): string => {
  const date = new Date(dateStr);
  const pad = (num: number) => num.toString().padStart(2, '0');
  
  return `${date.getFullYear()}-${pad(date.getMonth() + 1)}-${pad(date.getDate())}T${pad(date.getHours())}:${pad(date.getMinutes())}`;
}; 