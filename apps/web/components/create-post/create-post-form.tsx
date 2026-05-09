"use client";

import { type FormEvent, useEffect, useMemo, useRef, useState } from "react";
import Image from "next/image";
import { useRouter } from "next/navigation";
import { ImagePlus, X } from "lucide-react";
import { CommunityShareSection } from "@/components/create-post/community-share-section";
import { Field, inputClassName } from "@/components/create-post/field";
import { FormSection } from "@/components/create-post/form-section";
import { SegmentedControl } from "@/components/create-post/segmented-control";
import { Stepper } from "@/components/create-post/stepper";
import { ToggleRow } from "@/components/create-post/toggle-row";
import { SelectControl } from "@/components/forms/select-control";

const termLengthOptions = ["4 months", "8 months", "12 months"];
const termSeasonOptions = ["Fall", "Spring", "Winter"];
const maxListingImages = 5;
const apiBase = process.env.NEXT_PUBLIC_API_URL ?? "http://localhost:8080";

type CreatedPost = {
  id: string;
};

export function CreatePostForm() {
  const router = useRouter();
  const [title, setTitle] = useState("");
  const [address, setAddress] = useState("");
  const [price, setPrice] = useState(1000);
  const [propertyType, setPropertyType] = useState("Apartment");
  const [rooms, setRooms] = useState(1);
  const [roomsOccupied, setRoomsOccupied] = useState(0);
  const [bathrooms, setBathrooms] = useState(1);
  const [gender, setGender] = useState("Coed");
  const [description, setDescription] = useState("");
  const [content, setContent] = useState("");
  const [isSublet, setIsSublet] = useState(false);
  const [selectedTermLengths, setSelectedTermLengths] = useState<string[]>([]);
  const [selectedTermSeasons, setSelectedTermSeasons] = useState<string[]>([]);
  const [listingImages, setListingImages] = useState<File[]>([]);
  const [submitError, setSubmitError] = useState("");
  const [isSubmitting, setIsSubmitting] = useState(false);
  const submitInFlightRef = useRef(false);
  const imagePreviews = useMemo(
    () =>
      listingImages.map((image) => ({
        image,
        previewUrl: URL.createObjectURL(image),
      })),
    [listingImages],
  );

  useEffect(() => {
    return () => {
      imagePreviews.forEach(({ previewUrl }) => URL.revokeObjectURL(previewUrl));
    };
  }, [imagePreviews]);

  const toggleTermLength = (option: string) => {
    setSelectedTermLengths((current) =>
      current.includes(option)
        ? current.filter((value) => value !== option)
        : [...current, option],
    );
  };

  const toggleTermSeason = (option: string) => {
    setSelectedTermSeasons((current) =>
      current.includes(option)
        ? current.filter((value) => value !== option)
        : [...current, option],
    );
  };

  const handleImageUpload = (files: FileList | null) => {
    if (!files) {
      return;
    }

    setListingImages((current) =>
      [...current, ...Array.from(files)].slice(0, maxListingImages),
    );
  };

  const removeListingImage = (imageIndex: number) => {
    setListingImages((current) =>
      current.filter((_, index) => index !== imageIndex),
    );
  };

  const handleSubmit = async (event: FormEvent<HTMLFormElement>) => {
    event.preventDefault();
    if (submitInFlightRef.current) {
      return;
    }

    submitInFlightRef.current = true;
    setSubmitError("");

    const token = localStorage.getItem("authToken") ?? localStorage.getItem("token");
    if (!token) {
      setSubmitError("Please log in before creating a listing.");
      submitInFlightRef.current = false;
      return;
    }

    const normalizedTitle = title.trim();
    const normalizedContent = content.trim();
    if (normalizedTitle.length < 3) {
      setSubmitError("Listing title must be at least 3 characters.");
      submitInFlightRef.current = false;
      return;
    }
    if (normalizedContent.length < 10) {
      setSubmitError("Detailed content must be at least 10 characters.");
      submitInFlightRef.current = false;
      return;
    }
    if (roomsOccupied > rooms) {
      setSubmitError("Rooms occupied cannot exceed total rooms.");
      submitInFlightRef.current = false;
      return;
    }

    setIsSubmitting(true);

    try {
      const createResponse = await fetch(`${apiBase}/api/posts/create`, {
        method: "POST",
        headers: {
          Authorization: `Bearer ${token}`,
          "Content-Type": "application/json",
        },
        body: JSON.stringify({
          title: normalizedTitle,
          content: normalizedContent,
          address: address.trim(),
          isSublet,
          price,
          bedrooms: rooms,
          roomsOccupied,
          bathrooms,
          description: description.trim(),
          gender: gender === "Coed" ? undefined : gender.replace(" Only", ""),
          propertyType,
          term: selectedTermSeasons[0],
        }),
      });

      const createdPost = (await createResponse.json().catch(() => null)) as CreatedPost | null;
      if (!createResponse.ok || !createdPost?.id) {
        throw new Error(readErrorText(createdPost, "Could not create listing."));
      }

      await Promise.all(
        listingImages.map((image, index) => {
          const formData = new FormData();
          formData.append("image", image);
          formData.append("order", String(index));

          return fetch(
            `${apiBase}/api/pictures/upload?ownerType=post&ownerId=${createdPost.id}`,
            {
              method: "POST",
              headers: {
                Authorization: `Bearer ${token}`,
              },
              body: formData,
            },
          ).then(async (response) => {
            if (!response.ok) {
              const message = await response.text();
              throw new Error(message || `Could not upload ${image.name}.`);
            }
          });
        }),
      );

      router.push("/post-listings");
    } catch (error) {
      setSubmitError(error instanceof Error ? error.message : "Could not create listing.");
    } finally {
      submitInFlightRef.current = false;
      setIsSubmitting(false);
    }
  };

  return (
    <form className="space-y-10" onSubmit={handleSubmit}>
      <FormSection title="Essential Information">
        <div className="space-y-6">
          <Field label="Listing Title">
            <input
              className={inputClassName}
              placeholder="e.g. Spacious Studio near St. George Campus"
              value={title}
              onChange={(event) => setTitle(event.target.value)}
            />
          </Field>

          <Field label="Address">
            <input
              className={inputClassName}
              placeholder="Full address of the property"
              value={address}
              onChange={(event) => setAddress(event.target.value)}
            />
          </Field>

          <div className="grid gap-6 md:grid-cols-2">
            <Field label="Price per Month ($)">
              <input
                className={inputClassName}
                type="number"
                min={0}
                value={price}
                onChange={(event) => setPrice(Number(event.target.value))}
              />
            </Field>

            <Field label="Property Type">
              <SelectControl
                label="Property Type"
                options={["Apartment", "House", "Studio", "Shared Room", "Dorm"]}
                variant="create"
                value={propertyType}
                onChange={setPropertyType}
              />
            </Field>
          </div>
        </div>
      </FormSection>

      <FormSection title="Room & Unit Specs">
        <div className="space-y-8">
          <div className="grid gap-6 md:grid-cols-3">
            <Stepper label="Total Rooms" initialValue={1} min={1} value={rooms} onChange={setRooms} />
            <Stepper label="Rooms Occupied" initialValue={0} value={roomsOccupied} onChange={setRoomsOccupied} />
            <Stepper label="Bathrooms" initialValue={1} min={1} value={bathrooms} onChange={setBathrooms} />
          </div>

          <SegmentedControl
            label="Gender"
            options={["Female Only", "Male Only", "Coed"]}
            defaultValue="Coed"
            value={gender}
            onChange={setGender}
            size="sm"
          />

          <ToggleRow enabled={isSublet} onChange={setIsSublet} />
        </div>
      </FormSection>

      {isSublet && (
        <FormSection title="Term Details">
          <div className="space-y-8">
            <MultiSelectPills
              label="Term Season"
              options={termSeasonOptions}
              selectedOptions={selectedTermSeasons}
              onToggle={toggleTermSeason}
            />

            <MultiSelectPills
              label="Term Length"
              options={termLengthOptions}
              selectedOptions={selectedTermLengths}
              onToggle={toggleTermLength}
            />
          </div>
        </FormSection>
      )}

      <div className="space-y-7">
        <Field label="Quick Description (for card)">
          <textarea
            className={`${inputClassName} min-h-28 resize-y`}
            placeholder="Short summary of the housing..."
            value={description}
            onChange={(event) => setDescription(event.target.value)}
          />
        </Field>

        <Field label="Detailed Content">
          <textarea
            className={`${inputClassName} min-h-56 resize-y`}
            placeholder="Describe your property, rules, roommates, community vibes, etc..."
            value={content}
            onChange={(event) => setContent(event.target.value)}
          />
        </Field>

        <Field label="Photos">
          <div className="space-y-4">
            <label className="flex min-h-36 cursor-pointer flex-col items-center justify-center rounded-2xl border border-dashed border-orange-200 bg-orange-50/30 px-6 py-8 text-center transition-all hover:border-brand-orange hover:bg-orange-50/60">
              <ImagePlus className="mb-3 h-8 w-8 text-brand-orange" />
              <span className="text-base font-bold text-gray-700">
                Upload property photos
              </span>
              <span className="mt-1 text-sm font-medium text-gray-400">
                Add up to {maxListingImages} images
              </span>
              <input
                type="file"
                accept="image/*"
                multiple
                className="sr-only"
                onChange={(event) => handleImageUpload(event.target.files)}
              />
            </label>

            {imagePreviews.length > 0 && (
              <div className="grid gap-3 sm:grid-cols-2 lg:grid-cols-5">
                {imagePreviews.map(({ image, previewUrl }, index) => (
                  <div
                    key={`${image.name}-${index}`}
                    className="relative rounded-2xl border border-gray-100 bg-white p-3 shadow-sm"
                  >
                    <button
                      type="button"
                      onClick={() => removeListingImage(index)}
                      className="absolute right-1.5 top-1.5 rounded-full bg-white p-1 text-gray-400 shadow-md transition-colors hover:text-brand-orange"
                      aria-label={`Remove ${image.name}`}
                    >
                      <X size={16} />
                    </button>
                    <Image
                      src={previewUrl}
                      alt={image.name}
                      width={160}
                      height={160}
                      unoptimized
                      className="mb-2 aspect-square w-full rounded-xl object-cover"
                    />
                    <p className="truncate text-xs font-bold text-gray-600">
                      {image.name}
                    </p>
                  </div>
                ))}
              </div>
            )}
          </div>
        </Field>
      </div>

      <CommunityShareSection />

      <div className="mb-10 flex flex-col items-end gap-3">
        {submitError && (
          <p className="max-w-xl text-right text-sm font-bold text-red-600">{submitError}</p>
        )}
        <button
          type="submit"
          disabled={isSubmitting}
          className="rounded-2xl bg-brand-orange px-10 py-4 text-base font-black uppercase tracking-[0.18em] text-white shadow-xl shadow-brand-orange/25 transition-all hover:bg-orange-600 active:scale-[0.99] disabled:cursor-not-allowed disabled:bg-gray-300 disabled:shadow-none"
        >
          {isSubmitting ? "Submitting..." : "Submit"}
        </button>
      </div>
    </form>
  );
}

function readErrorText(data: unknown, fallback: string) {
  if (typeof data === "object" && data !== null && "error" in data && typeof data.error === "string") {
    return data.error;
  }

  return fallback;
}

interface MultiSelectPillsProps {
  label: string;
  options: string[];
  selectedOptions: string[];
  onToggle: (option: string) => void;
}

function MultiSelectPills({
  label,
  options,
  selectedOptions,
  onToggle,
}: MultiSelectPillsProps) {
  return (
    <div>
      <p className="mb-2 text-sm font-bold uppercase tracking-[0.16em] text-gray-400">
        {label}
      </p>
      <div
        className="grid rounded-2xl border border-gray-100 bg-white p-1 shadow-sm"
        style={{ gridTemplateColumns: `repeat(${options.length}, minmax(0, 1fr))` }}
      >
        {options.map((option) => {
          const isSelected = selectedOptions.includes(option);

          return (
            <button
              key={option}
              type="button"
              onClick={() => onToggle(option)}
              aria-pressed={isSelected}
              className={`rounded-xl px-4 py-3 text-sm font-bold uppercase transition-all ${
                isSelected
                  ? "bg-brand-orange text-white shadow-lg shadow-brand-orange/20"
                  : "text-gray-400 hover:text-brand-orange"
              }`}
            >
              {option}
            </button>
          );
        })}
      </div>
    </div>
  );
}
