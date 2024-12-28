import { useState } from 'react';

export type SortField = {
  field: string;
  label: string;
};

export type SortDirection = 'asc' | 'desc';

interface SortSelectProps {
  availableFields: SortField[];
  onChange: (sortFields: string[]) => void;
  className?: string;
}

export const SortSelect = ({ availableFields, onChange, className = '' }: SortSelectProps) => {
  const [selectedField, setSelectedField] = useState<string>('');
  const [direction, setDirection] = useState<SortDirection>('asc');

  const handleFieldChange = (field: string) => {
    setSelectedField(field);
    if (field) {
      onChange([`${field}:${direction}`]);
    } else {
      onChange([]);
    }
  };

  const handleDirectionChange = (newDirection: SortDirection) => {
    setDirection(newDirection);
    if (selectedField) {
      onChange([`${selectedField}:${newDirection}`]);
    }
  };

  return (
    <div className={`flex items-center gap-2 ${className}`}>
      <select
        value={selectedField}
        onChange={(e) => handleFieldChange(e.target.value)}
        className="rounded-xl border-gray-300 shadow-sm focus:border-primary-500 focus:ring-primary-500"
      >
        <option value="">Sort by</option>
        {availableFields.map((field) => (
          <option key={field.field} value={field.field}>
            {field.label}
          </option>
        ))}
      </select>
      <div className="flex rounded-xl border border-gray-300 overflow-hidden">
        <button
          type="button"
          onClick={() => handleDirectionChange('asc')}
          className={`px-3 py-2 text-sm ${
            direction === 'asc'
              ? 'bg-gray-900 text-white'
              : 'bg-white text-gray-700 hover:bg-gray-50'
          }`}
        >
          ↑ Asc
        </button>
        <button
          type="button"
          onClick={() => handleDirectionChange('desc')}
          className={`px-3 py-2 text-sm border-l border-gray-300 ${
            direction === 'desc'
              ? 'bg-gray-900 text-white'
              : 'bg-white text-gray-700 hover:bg-gray-50'
          }`}
        >
          ↓ Desc
        </button>
      </div>
    </div>
  );
}; 