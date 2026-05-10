const fieldClassName =
  "h-12 w-full rounded-2xl border border-gray-100 bg-gray-50 px-5 text-sm font-bold text-gray-800 shadow-inner shadow-gray-900/5 outline-none transition-all focus:border-brand-orange focus:bg-white focus:ring-4 focus:ring-brand-orange/10";

interface ProfileFieldProps {
  label: string;
  value: string;
  onChange?: (value: string) => void;
  type?: string;
  min?: number;
  max?: number;
  placeholder?: string;
}

export function ProfileField({
  label,
  value,
  onChange,
  type = "text",
  min,
  max,
  placeholder,
}: ProfileFieldProps) {
  return (
    <label className="block">
      <FieldLabel>{label}</FieldLabel>
      <input
        type={type}
        value={value}
        onChange={(event) => onChange?.(event.target.value)}
        min={min}
        max={max}
        placeholder={placeholder}
        className={fieldClassName}
      />
    </label>
  );
}

interface GenderSelectProps {
  value: string;
  onChange: (value: string) => void;
}

export function GenderSelect({ value, onChange }: GenderSelectProps) {
  return (
    <label className="block">
      <FieldLabel>Gender</FieldLabel>
      <select
        value={value}
        onChange={(event) => onChange(event.target.value)}
        className={`${fieldClassName} appearance-none`}
      >
        <option value="">Prefer not to say</option>
        <option value="Male">Male</option>
        <option value="Female">Female</option>
      </select>
    </label>
  );
}

export function FieldLabel({ children }: { children: string }) {
  return (
    <span className="mb-1.5 block text-[10px] font-bold uppercase tracking-[0.22em] text-gray-500">
      {children}
    </span>
  );
}

export { fieldClassName };
