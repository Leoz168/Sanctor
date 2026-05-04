import { AppShell } from "@/components/layout/app-shell";
import { ProfileFormCard } from "@/components/profile/profile-form-card";
import { ProfileSidebar } from "@/components/profile/profile-sidebar";

export default function ProfilePage() {
  return (
    <AppShell surface="cream">
      <div className="mx-auto grid max-w-7xl gap-8 px-4 py-6 sm:px-6 lg:grid-cols-[300px_1fr] lg:px-8 lg:py-8">
        <ProfileSidebar />

        <ProfileFormCard />
      </div>
    </AppShell>
  );
}
