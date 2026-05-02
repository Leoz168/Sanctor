import { LucideIcon } from "lucide-react";

interface FloatingActionButtonProps {
  icon: LucideIcon;
  children: string;
}

export function FloatingActionButton({ icon: Icon, children }: FloatingActionButtonProps) {
  return (
    <button className="fixed bottom-8 right-6 sm:right-8 flex items-center gap-2 px-5 py-4 bg-brand-orange text-white rounded-full font-bold shadow-xl shadow-brand-orange/30 hover:bg-orange-600 transition-all active:scale-95">
      <Icon className="w-5 h-5" />
      <span>{children}</span>
    </button>
  );
}
