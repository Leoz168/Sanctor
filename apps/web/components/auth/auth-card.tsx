import Link from "next/link";
import { ReactNode } from "react";
import { LucideIcon } from "lucide-react";
import { IconInput } from "@/components/forms/icon-input";

export interface AuthField {
  label: string;
  name: string;
  type: string;
  placeholder: string;
  icon: LucideIcon;
}

interface AuthCardProps {
  title: string;
  description: string;
  fields: AuthField[];
  submitLabel: string;
  footerText: string;
  footerLinkLabel: string;
  footerHref: string;
  children?: ReactNode;
}

export function AuthCard({
  title,
  description,
  fields,
  submitLabel,
  footerText,
  footerLinkLabel,
  footerHref,
  children,
}: AuthCardProps) {
  return (
    <div className="w-full max-w-[560px] rounded-[2rem] border border-white bg-white/90 p-7 shadow-2xl shadow-orange-900/10 backdrop-blur-md sm:p-10">
      <div className="mb-8">
        <h2 className="text-4xl font-bold tracking-tight text-gray-900">{title}</h2>
        <p className="mt-4 max-w-xl text-base font-medium leading-relaxed text-gray-500 sm:text-lg">
          {description}
        </p>
      </div>

      <form className="space-y-5">
        {fields.map((field) => (
          <label key={field.name} className="block">
            <span className="mb-2 block text-base font-bold text-gray-700 sm:text-lg">
              {field.label}
            </span>
            <IconInput icon={field.icon} name={field.name} type={field.type} placeholder={field.placeholder} />
          </label>
        ))}

        {children}

        <button className="w-full rounded-[1.25rem] bg-brand-orange px-5 py-4 text-lg font-bold text-white shadow-xl shadow-brand-orange/30 transition-all hover:bg-orange-600 active:scale-[0.99]">
          {submitLabel}
        </button>
      </form>

      <p className="mt-9 text-center text-base font-medium text-gray-500 sm:text-lg">
        {footerText}{" "}
        <Link href={footerHref} className="font-bold text-brand-orange hover:text-orange-600">
          {footerLinkLabel}
        </Link>
      </p>
    </div>
  );
}
