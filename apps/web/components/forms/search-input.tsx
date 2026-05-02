import { Search } from "lucide-react";

interface SearchInputProps {
  placeholder: string;
}

export function SearchInput({ placeholder }: SearchInputProps) {
  return (
    <div className="relative flex-1">
      <Search className="absolute left-5 top-1/2 -translate-y-1/2 w-5 h-5 text-gray-400" />
      <input
        type="text"
        placeholder={placeholder}
        className="w-full bg-brand-cream border border-gray-100 rounded-3xl pl-14 pr-6 py-4 text-base font-semibold focus:outline-none focus:bg-white focus:ring-2 focus:ring-brand-orange/20 transition-all shadow-inner"
      />
    </div>
  );
}
