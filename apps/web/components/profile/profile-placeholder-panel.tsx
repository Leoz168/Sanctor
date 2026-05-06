import type { ReactNode } from "react";
import { profilePanelClassName } from "@/components/profile/profile-panel-styles";

interface ProfilePlaceholderPanelProps {
  title: string;
  description: string;
  icon: ReactNode;
}

export function ProfilePlaceholderPanel({
  title,
  description,
  icon,
}: ProfilePlaceholderPanelProps) {
  return (
    <section className={profilePanelClassName}>
      <div className="flex h-14 w-14 items-center justify-center rounded-2xl bg-brand-cream text-brand-orange">
        {icon}
      </div>
      <h1 className="mt-6 text-3xl font-bold tracking-tight text-gray-900">
        {title}
      </h1>
      <p className="mt-3 max-w-2xl text-base font-semibold leading-7 text-gray-500">
        {description}
      </p>
    </section>
  );
}
