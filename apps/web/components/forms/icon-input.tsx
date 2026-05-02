import { InputHTMLAttributes } from "react";
import { LucideIcon } from "lucide-react";

interface IconInputProps extends InputHTMLAttributes<HTMLInputElement> {
  icon: LucideIcon;
}

export function IconInput({ icon: Icon, className = "", ...props }: IconInputProps) {
  return (
    <div className="relative">
      <Icon className="absolute left-5 top-1/2 h-5 w-5 -translate-y-1/2 text-gray-400" />
      <input
        {...props}
        className={`w-full rounded-[1.25rem] border border-gray-100 bg-brand-cream py-4 pl-14 pr-5 text-base font-semibold text-gray-900 outline-none transition-all placeholder:text-gray-400 focus:bg-white focus:ring-2 focus:ring-brand-orange/20 sm:text-lg ${className}`}
      />
    </div>
  );
}
