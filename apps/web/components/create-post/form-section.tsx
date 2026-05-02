import { ReactNode } from "react";

interface FormSectionProps {
  title: string;
  children: ReactNode;
}

export function FormSection({ title, children }: FormSectionProps) {
  return (
    <section className="rounded-[2rem] border border-gray-100 bg-white/85 p-6 shadow-sm sm:p-10">
      <div className="mb-8 border-b border-orange-100 pb-5">
        <h2 className="text-sm font-bold uppercase tracking-[0.28em] text-brand-orange/70">
          {title}
        </h2>
      </div>
      {children}
    </section>
  );
}
