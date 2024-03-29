/*
 * CAEP SSF API
 * CAEP.dev Receiver SSF API
 *
 * OpenAPI spec version: 1.0.0-oas3
 * Contact: hello@caep.dev
 *
 * NOTE: This class is auto generated by the swagger code generator program.
 * https://github.com/swagger-api/swagger-codegen.git
 * Do not edit the class manually.
 */

package io.swagger.client.model;

import java.util.Objects;
import java.util.Arrays;
import com.google.gson.TypeAdapter;
import com.google.gson.annotations.JsonAdapter;
import com.google.gson.annotations.SerializedName;
import com.google.gson.stream.JsonReader;
import com.google.gson.stream.JsonWriter;
import io.swagger.client.model.Delivery;
import io.swagger.v3.oas.annotations.media.Schema;
import java.io.IOException;
import java.util.ArrayList;
import java.util.List;
/**
 * CreateStreamResponse
 */

@javax.annotation.Generated(value = "io.swagger.codegen.v3.generators.java.JavaClientCodegen", date = "2023-12-07T14:46:31.358537509Z[GMT]")

public class CreateStreamResponse {
  @SerializedName("stream_id")
  private String streamId = null;

  @SerializedName("iss")
  private String iss = null;

  @SerializedName("aud")
  private String aud = null;

  @SerializedName("events_supported")
  private List<String> eventsSupported = null;

  @SerializedName("events_requested")
  private List<String> eventsRequested = null;

  @SerializedName("events_delivered")
  private List<String> eventsDelivered = null;

  @SerializedName("description")
  private String description = null;

  @SerializedName("delivery")
  private Delivery delivery = null;

  public CreateStreamResponse streamId(String streamId) {
    this.streamId = streamId;
    return this;
  }

   /**
   * Get streamId
   * @return streamId
  **/
  @Schema(description = "")
  public String getStreamId() {
    return streamId;
  }

  public void setStreamId(String streamId) {
    this.streamId = streamId;
  }

  public CreateStreamResponse iss(String iss) {
    this.iss = iss;
    return this;
  }

   /**
   * Get iss
   * @return iss
  **/
  @Schema(description = "")
  public String getIss() {
    return iss;
  }

  public void setIss(String iss) {
    this.iss = iss;
  }

  public CreateStreamResponse aud(String aud) {
    this.aud = aud;
    return this;
  }

   /**
   * Get aud
   * @return aud
  **/
  @Schema(description = "")
  public String getAud() {
    return aud;
  }

  public void setAud(String aud) {
    this.aud = aud;
  }

  public CreateStreamResponse eventsSupported(List<String> eventsSupported) {
    this.eventsSupported = eventsSupported;
    return this;
  }

  public CreateStreamResponse addEventsSupportedItem(String eventsSupportedItem) {
    if (this.eventsSupported == null) {
      this.eventsSupported = new ArrayList<String>();
    }
    this.eventsSupported.add(eventsSupportedItem);
    return this;
  }

   /**
   * Get eventsSupported
   * @return eventsSupported
  **/
  @Schema(description = "")
  public List<String> getEventsSupported() {
    return eventsSupported;
  }

  public void setEventsSupported(List<String> eventsSupported) {
    this.eventsSupported = eventsSupported;
  }

  public CreateStreamResponse eventsRequested(List<String> eventsRequested) {
    this.eventsRequested = eventsRequested;
    return this;
  }

  public CreateStreamResponse addEventsRequestedItem(String eventsRequestedItem) {
    if (this.eventsRequested == null) {
      this.eventsRequested = new ArrayList<String>();
    }
    this.eventsRequested.add(eventsRequestedItem);
    return this;
  }

   /**
   * Get eventsRequested
   * @return eventsRequested
  **/
  @Schema(description = "")
  public List<String> getEventsRequested() {
    return eventsRequested;
  }

  public void setEventsRequested(List<String> eventsRequested) {
    this.eventsRequested = eventsRequested;
  }

  public CreateStreamResponse eventsDelivered(List<String> eventsDelivered) {
    this.eventsDelivered = eventsDelivered;
    return this;
  }

  public CreateStreamResponse addEventsDeliveredItem(String eventsDeliveredItem) {
    if (this.eventsDelivered == null) {
      this.eventsDelivered = new ArrayList<String>();
    }
    this.eventsDelivered.add(eventsDeliveredItem);
    return this;
  }

   /**
   * Get eventsDelivered
   * @return eventsDelivered
  **/
  @Schema(description = "")
  public List<String> getEventsDelivered() {
    return eventsDelivered;
  }

  public void setEventsDelivered(List<String> eventsDelivered) {
    this.eventsDelivered = eventsDelivered;
  }

  public CreateStreamResponse description(String description) {
    this.description = description;
    return this;
  }

   /**
   * Get description
   * @return description
  **/
  @Schema(description = "")
  public String getDescription() {
    return description;
  }

  public void setDescription(String description) {
    this.description = description;
  }

  public CreateStreamResponse delivery(Delivery delivery) {
    this.delivery = delivery;
    return this;
  }

   /**
   * Get delivery
   * @return delivery
  **/
  @Schema(description = "")
  public Delivery getDelivery() {
    return delivery;
  }

  public void setDelivery(Delivery delivery) {
    this.delivery = delivery;
  }


  @Override
  public boolean equals(java.lang.Object o) {
    if (this == o) {
      return true;
    }
    if (o == null || getClass() != o.getClass()) {
      return false;
    }
    CreateStreamResponse createStreamResponse = (CreateStreamResponse) o;
    return Objects.equals(this.streamId, createStreamResponse.streamId) &&
        Objects.equals(this.iss, createStreamResponse.iss) &&
        Objects.equals(this.aud, createStreamResponse.aud) &&
        Objects.equals(this.eventsSupported, createStreamResponse.eventsSupported) &&
        Objects.equals(this.eventsRequested, createStreamResponse.eventsRequested) &&
        Objects.equals(this.eventsDelivered, createStreamResponse.eventsDelivered) &&
        Objects.equals(this.description, createStreamResponse.description) &&
        Objects.equals(this.delivery, createStreamResponse.delivery);
  }

  @Override
  public int hashCode() {
    return Objects.hash(streamId, iss, aud, eventsSupported, eventsRequested, eventsDelivered, description, delivery);
  }


  @Override
  public String toString() {
    StringBuilder sb = new StringBuilder();
    sb.append("class CreateStreamResponse {\n");
    
    sb.append("    streamId: ").append(toIndentedString(streamId)).append("\n");
    sb.append("    iss: ").append(toIndentedString(iss)).append("\n");
    sb.append("    aud: ").append(toIndentedString(aud)).append("\n");
    sb.append("    eventsSupported: ").append(toIndentedString(eventsSupported)).append("\n");
    sb.append("    eventsRequested: ").append(toIndentedString(eventsRequested)).append("\n");
    sb.append("    eventsDelivered: ").append(toIndentedString(eventsDelivered)).append("\n");
    sb.append("    description: ").append(toIndentedString(description)).append("\n");
    sb.append("    delivery: ").append(toIndentedString(delivery)).append("\n");
    sb.append("}");
    return sb.toString();
  }

  /**
   * Convert the given object to string with each line indented by 4 spaces
   * (except the first line).
   */
  private String toIndentedString(java.lang.Object o) {
    if (o == null) {
      return "null";
    }
    return o.toString().replace("\n", "\n    ");
  }

}
