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
  googleLabel?: string;
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
  googleLabel,
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

        {googleLabel && (
          <button
            type="button"
            className="mt-3 flex w-full items-center justify-center gap-6 rounded-[0.2rem] border border-gray-300 bg-white px-5 py-3 text-base font-medium text-[#5f6368] shadow-md transition-all hover:bg-gray-50 active:scale-[0.99] sm:text-lg"
          >
            <GoogleMark />
            {googleLabel}
          </button>
        )}
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

function GoogleMark() {
  return (
    <svg
      aria-hidden="true"
      className="h-6 w-6 shrink-0"
      viewBox="0 0 48 48"
    >
      <path
        fill="#EA4335"
        d="M24 9.5c3.54 0 6.71 1.22 9.21 3.6l6.85-6.85C35.9 2.38 30.47 0 24 0 14.62 0 6.51 5.38 2.56 13.22l7.98 6.19C12.43 13.72 17.74 9.5 24 9.5Z"
      />
      <path
        fill="#4285F4"
        d="M46.98 24.55c0-1.57-.15-3.09-.38-4.55H24v9.02h12.94c-.58 2.96-2.26 5.48-4.78 7.18l7.73 6c4.51-4.18 7.09-10.36 7.09-17.65Z"
      />
      <path
        fill="#FBBC05"
        d="M10.53 28.59A14.49 14.49 0 0 1 9.75 24c0-1.59.27-3.14.78-4.59l-7.98-6.19A23.94 23.94 0 0 0 0 24c0 3.88.93 7.56 2.56 10.78l7.97-6.19Z"
      />
      <path
        fill="#34A853"
        d="M24 48c6.48 0 11.93-2.13 15.89-5.81l-7.73-6c-2.15 1.45-4.92 2.3-8.16 2.3-6.26 0-11.57-4.22-13.47-9.91l-7.98 6.19C6.51 42.62 14.62 48 24 48Z"
      />
    </svg>
  );
}
