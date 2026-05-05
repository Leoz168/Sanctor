const fieldClassName =
  "h-12 w-full rounded-2xl border border-gray-100 bg-gray-50 px-5 text-sm font-bold text-gray-800 shadow-inner shadow-gray-900/5 outline-none transition-all focus:border-brand-orange focus:bg-white focus:ring-4 focus:ring-brand-orange/10";

interface ProfileFieldProps {
  label: string;
  value: string;
  type?: string;
  min?: number;
  max?: number;
}

export function ProfileField({
  label,
  value,
  type = "text",
  min,
  max,
}: ProfileFieldProps) {
  return (
    <label className="block">
      <FieldLabel>{label}</FieldLabel>
      <input
        type={type}
        defaultValue={value}
        min={min}
        max={max}
        className={fieldClassName}
      />
    </label>
  );
}

export function GenderSelect() {
  return (
    <label className="block">
      <FieldLabel>Gender</FieldLabel>
      <select defaultValue="Male" className={`${fieldClassName} appearance-none`}>
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
