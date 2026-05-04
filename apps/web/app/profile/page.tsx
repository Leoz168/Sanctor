import { AppShell } from "@/components/layout/app-shell";
import { ProfileActivityCard } from "@/components/profile/profile-activity-card";
import { ProfileFormCard } from "@/components/profile/profile-form-card";
import { ProfileSidebar } from "@/components/profile/profile-sidebar";

export default function ProfilePage() {
  return (
    <AppShell surface="cream">
      <div className="mx-auto grid max-w-7xl gap-8 px-4 py-10 sm:px-6 lg:grid-cols-[300px_1fr] lg:px-8 lg:py-14">
        <ProfileSidebar />

        <div className="space-y-6">
          <ProfileFormCard />
          <ProfileActivityCard />
        </div>
      </div>
    </AppShell>
  );
}
